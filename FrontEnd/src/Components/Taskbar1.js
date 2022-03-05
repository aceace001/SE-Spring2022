import React from 'react';
import {Link} from "react-router-dom";

const TaskBar1 = () => {
    return (
        <div className={"d-flex flex-column flex-shrink-0 p-3 text-white bg-dark"} >
            <ul className={"nav-pills flex-column mb-auto mb-4 p-2"}>
                <li className={"nav-item"}>
                    <Link to="/" className="nav-link px-2 text-white">Home</Link>
                </li>
                <li className={"nav-item"}>
                    <Link to="/HomePost" className="nav-link px-2 text-white">Post Page</Link>
                </li>
            </ul>
        </div>
    );
};

export default TaskBar1;