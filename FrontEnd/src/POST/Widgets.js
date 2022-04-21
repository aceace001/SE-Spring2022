import React from "react";
import "./Widgets.css";
import {

    TwitterTweetEmbed,
} from "react-twitter-embed";
import SearchIcon from "@material-ui/icons/Search";

function Widgets() {
    return (
        <div className="widgets">


            <div className="widgets__widgetContainer">
                <h2>What's happening</h2>

                <TwitterTweetEmbed tweetId={"1507437682636562433"} />

            </div>
        </div>
    );
}

export default Widgets;