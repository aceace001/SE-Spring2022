import React, {useState} from 'react';
import TaskBar1 from "../TaskBar/TaskBar1";

const HomePost = () => {
    const [name, setName] = useState('');
    return (
        <TaskBar1 name={name} setName={setName}/>
    );
};

export default HomePost;