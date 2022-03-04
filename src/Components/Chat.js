import "./Chat.css"

export default function Chat({own}) {
    return (
        <div className={own ? "TextOwn" : "Texts"}>
             <div className="ChatTop">
                 <img className="ChatImg"
                      src={require("./profile.png")}  alt=""/>
                 <p className="Text">
                     Hello! AAAAAAAAAAAAA
                 </p>
             </div>
            <div className="ChatBottom">
                2 hours ago
            </div>


        </div>
    )


}