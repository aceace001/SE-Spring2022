import React from 'react';

const showFriends = ({showfriends}) => {
    return (
        <>
        {showfriends.map((val,key) =>{
        return(
        <>
        <div className="showFriends">
            <p>{val.first_name}</p>
        </div>
        </>
        )
        })}
        </>
    )
}


export default showFriends;
