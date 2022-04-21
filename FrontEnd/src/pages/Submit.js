import React, {SyntheticEvent, useState} from 'react';
import {Navigate} from 'react-router-dom';

const Register = () => {
    const [Author, setName] = useState('');
    const [Title, setTitle] = useState('');
    const [navi, setNavi] = useState(false);

    console.log({
        Author,
        Title

    })

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();
        const response = await fetch('http://localhost:8081/post', {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                Author,
                Title

            })
        })
        setNavi(true);


        const content = await response.json();

        console.log(content);
    }

    if (navi) {
        return <Navigate to="/HomePost1"/>;
    }

    return (
        <div className="register">
            <form className={"p-4 p-md-5 border rounded-3 bg-light"} onSubmit={submit}>
                <h1 className="h3 mb-3 fw-normal">What's happening</h1>
                <div className={"form-floating mb-3"} id='first'>
                    <input type="username" className="form-control" placeholder="Enter your username to post!" required
                           onChange={e => setName(e.target.value)}
                    />
                    <label>Enter your username to post!</label>
                </div>

                <div className={"form-floating mb-3"} id='last'>
                    <input type="username" className="form-control" placeholder="Enter your tweets" required
                           onChange={e => setTitle(e.target.value)}
                    />
                    <label>Enter the tweet you want to post</label>
                </div>

                <button className="w-100 btn btn-lg btn-primary" type="submit">Post Now!</button>
            </form>
        </div>
    );
};

export default Register;
