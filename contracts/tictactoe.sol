pragma solidity ^0.4.15;

/*
interface Protochannel {
    *//*
       @dev Apply operation enclosed within state and return new state
       @param state is a minivm memory
       @param v, v part of the personalSignature of state
       @param r, r part of the personalSignature of state
       @param s, s part of the personalSignature of state
    *//*
    function Apply(bytes state, uint8 v, bytes32 r, bytes32 s) public view returns (bytes);
    function ValidateCheckpoint(bytes state, uint8 v, bytes32 r, bytes32 s) public;
    function Challenge(bytes state, byte[3] newOperation, uint8[] v, bytes32[] r, bytes32[] s) public;
    function Finalize() public;
}
*/
contract TicTacToe /* is Protochannel */ {

    // Position
    uint public x_max;
    uint public y_max;

    uint constant hdr_length = 4;
    uint constant timeout = 30; // 30 s

    address public Judge;

    address public P1;
    address public P2;

    address public FinalizeTo;
    uint public FinalizeTimeout;
    byte public FinalizeNonce;

    event NewGame(address player1, address player2);
    event Challenged(address player, uint timeout);
    event Settle(address player);

    enum Operation { OP_NOP, OP_PUT }

    modifier checkMap(uint pos) {
        require(pos < (y_max * x_max));
        _;
    }

    modifier isPlayer() {
        require(P1 == msg.sender ||
                P2 == msg.sender ||
                Judge == msg.sender);
        _;
    }

    modifier isJudge() {
        require(Judge == msg.sender);
        _;
    }

    modifier isPlayable() {
        require(P1 != 0x0);
        require(P2 != 0x0);
        _;
    }

    function TicTacToe() public {
        Judge = msg.sender;
        y_max = 3;
        x_max = 3;
        P1 = 0x533A245F03A1a46CAcB933a3bEEF752fD8ff45C3;
        P2 = 0x48d24715FE9F7daa90286141f9b17e184B5A148B;
        FinalizeTimeout = 0;
        FinalizeNonce = 0x0;
        FinalizeTo = 0;
    }

    function SetPlayer(address other) public payable {
        require(P1 == 0x0 && P2 == 0x0);
        require(other != msg.sender); // no solo game !
        require(other != 0x0);
        P1 = msg.sender;
        P2 = other;
        NewGame(P1, P2);
    }

    function GetPlayerMark(address player) public view returns (byte) {
        if (P1 == player) return 1;
        if (P2 == player) return 2;
        return 0;
    }

    function InfoSigner(bytes message, uint8 v, bytes32 r, bytes32 s) public pure returns (address) {
        return getSigner(message, v, r, s);
    }

    function getSigner(bytes message, uint8 v, bytes32 r, bytes32 s) internal pure returns (address) {
        bytes memory prefix = "\x19Ethereum Signed Message:\n13"; // TODO dynamic state length
        bytes32 prefixedHash = keccak256(prefix, message);
        return ecrecover(prefixedHash, v, r, s);
    }

    function CreateBoard() public pure returns (bytes) {
        /*
           Game is a minivm memory
            byte 0 next operation (NOP: 0; PUT([0; x_max*Y+X[);
            byte 1 operation counter
            byte 2,3 params binary encoded params for operation
            byte 4-  game board
           Board should be modified if operation can be applied
        */
        bytes memory gamestate = new bytes(13);
        gamestate[0]  = byte(uint(Operation.OP_NOP)); /* next operation*/
        gamestate[1]  = byte(0); /* operation counter */
        gamestate[2]  = byte(0); /* 1st param */
        gamestate[3]  = byte(0); /* 2nd param */
        gamestate[4]  = byte(0);
        gamestate[5]  = byte(0);
        gamestate[6]  = byte(0);
        gamestate[7]  = byte(0);
        gamestate[8]  = byte(0);
        gamestate[9]  = byte(0);
        gamestate[10] = byte(0);
        gamestate[11] = byte(0);
        gamestate[12] = byte(0);
        return gamestate;
    }

    function CheckAndApply(bytes state, uint8 v, bytes32 r, bytes32 s) public view returns (bytes) {
        require(Verify(state, v, r, s));
        return Apply(state, getSigner(state, v, r, s));
    }

    //apply operation on state
    function Apply(bytes state, address player) public view returns (bytes) {
        Operation operation = Operation(uint(state[0]));
        if (operation == Operation.OP_PUT) {
            uint x = uint(state[2]);
            uint y = uint(state[3]);
            state[hdr_length + x_max*y+x] = GetPlayerMark(player);
        }
        return state;
    }

    function can_put(bytes state) public view returns (bool) {
        uint x = uint(state[2]);
        uint y = uint(state[3]);
        return state[hdr_length + x_max*y+x] == 0;
    }

    //encode put
    function op_put(bytes state, uint x, uint y) public pure returns (bytes) {
        state[0] = byte(uint(Operation.OP_PUT)); // PUT
        state[1] = byte(uint(state[1])+1);   // operation counter
        state[2] = byte(x);
        state[3] = byte(y);
        return state;
    }

    /*
       @dev Verify an operation on state
       @param State is a minivm memory
        byte 0 current operation (NOP: 0; PUT([0; x_max*Y+X[);
        byte 1 operation counter
        byte 2,3 params binary encoded params for operation
        byte 4-  state board
       Board should be modified if operation can be applied
    */
    function Verify(bytes state, uint8 v, bytes32 r, bytes32 s) public constant returns (bool) {
        // check board is signed
        address current_player = getSigner(state, v, r, s);
        if (state.length <= hdr_length) return false;
        Operation operation = Operation(uint(state[0]));
        uint8 counter = uint8(state[1]);
        // NOP
        // check if last operation + new operation on board match new board
        if (operation == Operation.OP_NOP) {
            // TODO check if map empty
            return current_player == Judge;
        }
        if (P1 != current_player && P2 != current_player) return false;
        // Can't play if not your turn
        if ((current_player == P1 && (counter&0x1) == 0) || // Pair turn is Player1
            (current_player == P2 && (counter&0x1) == 1)) return false; // Odd turn is Player2
        // PUT(pos)
        if (operation == Operation.OP_PUT) {
            // check if operation can apply
            return can_put(state);
        }
        return false;
    }

    function checkY(bytes state, uint y, byte m) public view returns (bool) {
        return state[hdr_length + (x_max*y)] == state[hdr_length + (x_max*y)+1] &&
               state[hdr_length + (x_max*y)+1] == state[hdr_length + (x_max*y)+2] &&
               m == state[hdr_length + (x_max*y)+2];
    }

    function checkX(bytes state, uint x, byte m) public view returns (bool) {
        return state[hdr_length + (x_max*0)+x] == state[hdr_length + (x_max*1)+x] &&
               state[hdr_length + (x_max*1)+x] == state[hdr_length + (x_max*2)+x] &&
               m == state[hdr_length + (x_max*2)+x];
    }

    function checkDiag(bytes state, byte m) public view returns (bool) {
        return state[hdr_length + (x_max*0)+0] == state[hdr_length + (x_max*1)+1] &&
               state[hdr_length + (x_max*1)+1] == state[hdr_length + (x_max*2)+2] &&
               state[hdr_length + (x_max*0)+2] == state[hdr_length + (x_max*1)+1] &&
               state[hdr_length + (x_max*1)+1] == state[hdr_length + (x_max*2)+0] &&
               m == state[hdr_length + (x_max*1)+1];
    }

    function checkLines(bytes state, address player) public view returns (bool) {
      byte m = GetPlayerMark(player);
      return checkY(state, 0, m) ||
             checkY(state, 1, m) ||
             checkY(state, 2, m) ||
             checkX(state, 0, m) ||
             checkX(state, 1, m) ||
             checkX(state, 2, m) ||
             checkDiag(state, m);
    }

    function getWinner(bytes state) public view returns (address, bool) {
        if (checkLines(state, P1))
            return (P1, true);
        if (checkLines(state, P2))
            return (P2, true);
        uint8 counter = uint8(state[1]);
        // Nonce = 0, no winner
        if (counter == 0)
            return (0x0, false);
        // No winner according to rule => winner is determined according to state determined
        // TODO TEST
        return (((counter % 2) == 0 ? P1 : P2), true);
    }

    function setNewOperation(bytes memory state, byte[3] newOperation) private pure returns (bytes) {
        state[0] = newOperation[0];
        state[2] = newOperation[1];
        state[3] = newOperation[2];
        return state;
    }

    function Challenge(bytes state, byte[3] newOperation, uint8[] v, bytes32[] r, bytes32[] s) isPlayer public {
      address sig1 = getSigner(state, v[0], r[0], s[0]);
      address sig2 = msg.sender;
      require((P1 == sig1 && P2 == sig2) || (P1 == sig2 && P2 == sig1));
      require(FinalizeNonce < state[1]); // Nonce can only go forward

      address winner;
      bool ok;
      bytes memory lastState = setNewOperation(state, newOperation); // Encode new operation on state
      lastState = Apply(lastState, P2); // Apply new operation on state
      // TODO check
      (winner, ok) = getWinner(lastState); // Determine winner
      require(ok);

      FinalizeTo = winner;
      FinalizeNonce = state[1]; // nonce
      FinalizeTimeout = now + timeout;
      Challenged(winner, FinalizeTimeout); // state ?
    }

    function Finalize() isPlayer public {
        require(FinalizeTo != 0x0);
        require(FinalizeTimeout < now);
        Settle(FinalizeTo);
        selfdestruct(FinalizeTo);
    }
}
