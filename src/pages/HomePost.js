import React, {useState, useEffect} from 'react';
import TaskBar1 from "../Components/Taskbar1";
const HomePost = () => {
    const [name, setName] = useState('');


    useEffect(() => {
        (
            async () => {
                const response = await fetch('http://localhost:8081/api/user', {
                    headers: {'Content-Type': 'application/json'},
                    credentials: 'include',
                });

                const content = await response.json();

                setName(content.name);
            }
        )();
    });

    return (
        <div>
            {name ? "Welcome" + name :"Wrong username or password"}
        </div>
    )
}
export default HomePost;
