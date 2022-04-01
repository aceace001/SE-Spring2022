import React, { Component } from 'react';
import './ChatInput.css';

class ChatInput extends Component {

    render() {
        return (

                <div className='ChatInput'>
                    <input type="text" onKeyDown={this.props.send}  placeholder="Type a message and Enter to send"/>
                </div>


        );
    };

}

export default ChatInput;
