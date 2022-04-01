import React, { Component } from 'react';
import './ChatHistory.css';
import Message from '../Messages1';

class ChatHistory extends Component {
    render() {

        const messages = this.props.chatHistory.map(msg => <Message key={msg.timeStamp} message={msg.data} />);

        return (
            <div className='ChatHistory'>
                {messages}
            </div>
        );
    };

}

export default ChatHistory;

