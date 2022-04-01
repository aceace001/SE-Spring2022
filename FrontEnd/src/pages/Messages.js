import React, { Component } from 'react';
import ChatHistory from '../Components/Chats/ChatHistory';
import ChatInput from '../Components/ChatInput/ChatInput';
import { connect, sendMsg } from '../api';
import Friends from "../Components/Friends";
import Online from "../Components/Online";


class MessageMain extends Component {
    constructor(props) {
        super(props);
        this.state = {
            chatHistory: []
        }
    }

    componentDidMount() {
        connect((msg) => {
            console.log("New Message")
            this.setState(prevState => ({
                chatHistory: [...prevState.chatHistory, msg]
            }))
            console.log(this.state);
        });
    }

    send(event) {
        if (event.keyCode === 13) {
            sendMsg(event.target.value);
            event.target.value = "";
        }
    }

    render() {
        return (
            <div className="App">

                <>
                    <div className="messenger">
                        <div className="chatMenu">
                            <div className="chatMenuWrapper">
                                <input placeholder="Search for friends" className="chatMenuInput"/>
                                <Friends/>
                                <Friends/>
                                <Friends/>
                                <Friends/>
                                <Friends/>

                            </div>
                        </div>

                        <div className="chatBox">
                            <div className="currentchatter">
                                <b>Arthur</b>
                            </div>

                            <ChatHistory chatHistory={this.state.chatHistory} />

                            <ChatInput send={this.send} />

                        </div>
                        <div className="chatOnline">
                            <div className="chatOnlineWrapper">
                                <Online/>
                                <Online/>
                                <Online/>
                                <Online/>
                                <Online/>
                            </div>
                        </div>
                    </div>
                </>
                );
                }

            </div>
        );
    }
}

export default MessageMain;
