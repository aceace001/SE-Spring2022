import React, {useState, useEffect} from 'react';
import TaskBar1 from "../Components/Taskbar1";
const HomePost = () => {
    const [email, setEmail] = useState('');

    const [password, setPassword] = useState('');


    useEffect(() => {
        (
            async () => {
                const response = await fetch('http://localhost:8081/api/user', {
                    headers: {'Content-Type': 'application/json'},
                    credentials: 'include',
                });

                const content = await response.json();

                setEmail(content.email);
                setPassword(content.password)
            }
        )();
    });

    return (
        <div>
            {email ? "Welcome " + email :"Wrong username or password"}
        </div>
    )
}
export default HomePost;
