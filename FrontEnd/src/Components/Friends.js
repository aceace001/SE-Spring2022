import "./Friends.css"

export default function Friends() {
    return (
        <div className = "Friends">
            <img className="FriendImg"
                 src={require("./profile.png")}  alt=""/>
            <span className="FriendsName"><b>Arthur</b></span>
        </div>
    )
}