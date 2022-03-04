import "./Online.css";

export default function Online() {

    return (
        <li className="OnlineFriend">
            <div className="OnlineProfile">
                <img className="OnlineProfileImg" src={require("./profile.png")}  alt=""/>
                <span className="Status"></span>
            </div>
            <span className="OnlineUser"><b>Arthur</b></span>
        </li>
    );
}