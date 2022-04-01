import React, { useState } from 'react';

import { styles } from '../styles'

import { LoadingOutlined } from '@ant-design/icons'
import Avatar from '../Avatar';

import axios from 'axios';

const EmailForm = props => {
    const [email, setEmail] = useState('')
    const [loading, setLoading] = useState(false)
    //we need to have something like restAPI here
    function getOrCreateUser(callback){
        axios.put(
            'https://api.chatengine.io/users/',
            {
                'username': email,
                'secret': email,
                'email': email,
            },
            {headers: {"Private-Key": 'cddcaddc-75fd-4618-bd05-33e619f719a4'}}
        )
        .then(r => callback(r.data))
        .catch(e => console.log('Get or create user error', e))
    }

    function getOrCreateChat(callback){
        axios.put(
            'https://api.chatengine.io/chats/',
            {
                'username': {"Fugang Deng": email},
                'is_direct_caht': true
            },
            {headers: {"Project-ID": '298755fc-0e1a-4390-ab40-7134b4727cc0', "User-Name": email, "User-Secret": email}}
        )
        .then(r => callback(r.data))
        .catch(e => console.log('Get or create chat error', e))
    }

    function handleSubmit(event){
        event.preventDefault();
        setLoading(true)
        console.log('Sending email', email)

        //Code to handle restAPI here
        getOrCreateUser(
            user => {
                // getOrCreateChat(
                //     chat => console.log('success',chat)
                // )
                props.setUser && props.setUser(user)
                getOrCreateChat(
                    chat => {
                        setLoading(false)
                        props.setChat && props.setChat(chat)
                    })
            }
        )
        
    }
    return(
        <div
            style={{
                ...styles.emailFormWindow,
                ...{
                    height: props.visible ? '100%': '0px',
                    opacity: props.visible ? '1' : '0',
                }
            }}
        >
           
            <div style={{ heigth: '0px' }}>  
                <div style={styles.stripe} />
            </div>
            <div
                className='transition-5'
                style={{ 
                    ...styles.loadingDiv,
                    ...{
                        zIndex: loading ? '10' : '-1',
                        opacity: loading ? '0.33' : '0',
                    }
                }}
            />
            <LoadingOutlined
                className='transition-5'
                style={{
                    ...styles.loadingIcon,
                    ...{
                        zIndex: loading ? '10':'-1',
                        opacity: loading ? '1':'0',
                        fontSize: '82px', //The calc below is help to stay in the middle
                        top: 'calc(50% - 41px)', //82/2 = 41px
                        left: 'calc(50% - 41px)',
                    } 
                }}
            />
            <div style={{ position: 'absolute', height: '100%', width: '100%', textAlign: 'center'}}>
                <Avatar
                    style={{
                        position: 'relative',
                        left: 'calc(50% - 44px)',
                        top: '-55%',
                    }}
                />
                <div style={styles.topText}>
                    Welcome to my <br/> support

                </div>
                <form 
                    onSubmit={e => handleSubmit(e)}
                    style={{ position: 'relative', width: '100%', top: '-40%'}}
        
                >
                    <input
                        style={styles.emailInput}
                        onChange={e => setEmail(e.target.value)}
                        placeholder='Your Email'
                    />
                </form>
                <div style={styles.bottomText}>
                    Enter your email <br /> to get supports
                </div>
            </div>
        </div>
    )
}

export default EmailForm