import React, { useState, useEffect } from 'react'
import { ChatEngineWrapper, Socket, ChatFeed} from 'react-chat-engine'


const ChatEngine = props => {
    const [showChat, setShowChat] = useState(false)

    useEffect(() => {
        if (props.visible){
            setTimeout(() => {
                setShowChat(true)
            }, 500)
        }
    })
    return(
        <div
            className="transition-5"
            style={{
                
                height: props.visible ? '100%' : '0px',
                zIndex: props.visible ? '100' : '0',
                width: '100%',
                backgroundColor: 'white',
            }}
        >
            {
                showChat &&
                <ChatEngineWrapper>
                    <Socket
                        projectID={'298755fc-0e1a-4390-ab40-7134b4727cc0'}
                        userName={props.user.email}
                        userSecret={props.user.email}
                    />
                    <ChatFeed activeChat={props.chat.id} />
                </ChatEngineWrapper>
            }
        </div>
    )
}

export default ChatEngine