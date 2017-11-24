pragma solidity ^0.4.15;

/*
contract Protochannel {
    address public owner;
    function Protochannel() {
        owner = msg.sender;
    }
}*/

contract TicTacToe /*is Protochannel*/ {

    // Position
    uint public x_max;
    uint public y_max;

    uint hdr_length = 4;

    address public Judge;

    address public P1;
    address public P2;

    event NewGame(address player1, address player2);

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
        P1 = 0x533a245f03a1a46cacb933a3beef752fd8ff45c3;
        P2 = 0x48d24715FE9F7daa90286141f9b17e184B5A148B;
    }

    function SetPlayer(address other) public payable {
        require(P1 == 0x0 && P2 == 0x0);
        require(other != msg.sender); // no solo game !
        require(other != 0x0);
        P1 = msg.sender;
        P2 = other;
        NewGame(P1, P2);
    }

    function GetPlayerMark(address player) public constant returns (byte) {
        if (P1 == player) return 1;
        if (P2 == player) return 2;
        return 0;
    }

    function getSigner(bytes message, uint8 v, bytes32 r, bytes32 s) internal pure returns (address) {
        bytes memory prefix = "\x19Ethereum Signed Message:\n13"; // TODO states length
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
        gamestate[0]  = byte(0); /* next operation*/
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

    function CheckAndApply(bytes state, uint8 v, bytes32 r, bytes32 s) public constant returns (bytes) {
        require(Verify(state, v, r, s));
        return Apply(state, getSigner(state, v, r, s));
    }
//apply operation on state
    function Apply(bytes state, address player) public constant returns (bytes) {
        byte operation = state[0];
        if (operation == 1) {
            uint x = uint(state[2]);
            uint y = uint(state[3]);
            state[hdr_length + x_max*y+x] = GetPlayerMark(player);
        }
        return state;
    }

    function can_put(bytes state) public constant returns (bool) {
        uint x = uint(state[2]);
        uint y = uint(state[3]);
        return state[hdr_length + x_max*y+x] == 0;
    }
//encode put
    function op_put(bytes state, uint x, uint y) public pure returns (bytes) {
        state[0] = 0x01; // PUT
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
        byte operation = state[0];
        uint8 counter = uint8(state[1]);
        // NOP
// check if last operation + new operation on board match new board
        if (operation == 0) {
            // TODO check if map empty
            return current_player == Judge;
        }
        if (P1 != current_player && P2 != current_player) return false;
        // Can't play if not your turn
        if ((current_player == P1 && (counter&0x1) == 0) ||
            (current_player == P2 && (counter&0x1) == 1)) return false;
        // PUT(pos)
        if (operation == 1) {
            // check op_code can apply
            return can_put(state);
        }
        return false;
    }

    //function Challenge(bytes state, bytes[] v, bytes32[] r, bytes32[] s) public isPlayer {
        // check state is signed
        //address submitter = getSigner(state, v, r, s);
    //}
}
