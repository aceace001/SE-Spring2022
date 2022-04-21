import React, {useState, useEffect} from 'react';
import Widgets from "../POST/Widgets"
import "./HomePost1.css"
const HomePost1 = () => {
    const [author, setName] = useState([]);

    //const [title, setTitle] = useState([]);


    useEffect(() => {
        (
            async () => {
                const response = await fetch('http://localhost:8081/post', {
                    headers: {'Content-Type': 'application/json'},
                    credentials: 'include',
                });

                const content = await response.json();

                setName(content.messges)
            }
        )();
    });

    return (
        <div className = "P">
        <div className="Posts">
            <b>Posts</b>
                {author.map((authors, i) => {
                    return (<div key={i}>{authors.author + ": " }
                        {authors.title}</div>)
                })}

        </div>

        <div className="Twitter">
            <b>Feed</b>
            <Widgets />
        </div>
        </div>
    )
}
export default HomePost1;
