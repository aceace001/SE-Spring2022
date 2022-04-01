import React, { useState } from 'react';
import {  styles } from './styles'

const Avatar = props => {
    const [hovered, setHovered] = useState(false)
    return(
        <div style={props.style}>
            <div
                className='transition-3'
                style={{
                    ...styles.avatarHello,
                    ...{ opacity: hovered ? '1':'0'}
                }}
            >
                Support Chat Here!!!
            </div>
            <div
                className='transition-3'
                onMouseEnter={() => setHovered(true)}
                onMouseLeave={() => setHovered(false)}
                onClick={() => props.onClick && props.onClick()}
                style={{     //double {} means merge style here
                    ...styles.chatWithMeButton,
                    ...{ border: hovered ? '1px solid #f9f0ff': '8px solid #ADD8E6'}
                }}
            />
        </div>
    )
}

export default Avatar