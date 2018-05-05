pragma solidity ^0.4.20;

contract MyContract {

    // owner contract
    address public owner;

    // map address user
    address[] public members;

    // total link
    uint256 public totalLink;

    // 1ETH = number get link (exp: exchange=100, send 1ETH to address, link will be generate 100 )
    uint256 public exchange;

    struct Link{
        uint256 id;
        address userManager; // address uplink
        address prefAddress; // address get link from
        string originLink; // link
        uint256 investEth; // number ETH invest to share link
        uint256 originId;     // id origin
    }

    // map content contain link and id (uint256)
    mapping(uint256 => Link) mapLinks;

    // mỗi địa chỉ sẽ quản lý một map Link riêng của mình
    mapping(address => mapping(uint256 => Link)) userMapLink;

    // number shared of link
    mapping(uint256 => uint256) countLink;

    // init
    function MyContract(uint256 _exchange) public {
        totalLink = 0;
        owner = msg.sender;
        if(_exchange != 0 ) exchange = _exchange;
    }

    function checkUserExists(address user) public constant returns (bool){
        for (uint256 i = 0; i < members.length; i++){
            if (members[i] == user) return true;
        }
        return false;
    }

    function checkLinkExists(uint256 linkId) public constant returns (bool){
        if (mapLinks[linkId].id == 0){
            return false;
        }
        return true;
    }

    function checkNumberGenerateLink(uint256 linkId) public constant returns (bool){
        Link storage link = mapLinks[linkId];
        uint256 totalShare = link.investEth*exchange;
        if (countLink[linkId] < totalShare){
            return true;
        }else{
            return false;
        }
    }

    function checkSenderShare(uint256 linkId, address sender) public constant returns (bool){
        if (userMapLink[sender][linkId].id == 0) return true; // sender chưa từng share link này
        return false;
    }

    function upLoadLink(string _content, uint256 _investEth) public payable{

        bool exist = checkUserExists(msg.sender);

        if (!exist) members.push(msg.sender);

        totalLink ++ ;

        var newLink = Link(
            totalLink,
            msg.sender,
            msg.sender,
            _content,
            _investEth,
            totalLink
        );

        mapLinks[totalLink] = newLink;
        userMapLink[msg.sender][totalLink] = newLink;
        countLink[totalLink] = 0;
        /*
            minus token
        */

        emit UpLoadLink(msg.sender, newLink);
    }

    function sendEther(address _receive, uint256 _numEther) public {
        _receive.transfer(_numEther);
    }

    function generateLink(uint256 linkId) public payable returns (bool){
        require(checkLinkExists(linkId));
        require(checkNumberGenerateLink(linkId));
        require(checkSenderShare(linkId, msg.sender));
        /*
            create new link for address
        */
        // get pref link by linkId
        Link storage link = mapLinks[linkId];
        Link memory newLink = Link(totalLink, msg.sender, link.userManager, link.originLink, link.investEth, link.originId);

        totalLink ++ ;

        //
        userMapLink[msg.sender][totalLink] = newLink;
        emit GenerateLink(msg.sender, newLink);
        sendEther(msg.sender, 10**16);
        emit SentEther(msg.sender, 10**16);
        return true;
    }

    event SentEther(address _sender, uint256 eth);
    event UpLoadLink(address _from, Link content);
    event GenerateLink(address _from, Link content);
}
