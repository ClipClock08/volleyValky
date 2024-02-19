import {Link, Outlet, useNavigate} from "react-router-dom";
import {useEffect, useState} from "react";
import Alert from "./components/Alert";

function App() {
    const [jwtToken, setJwtToken] = useState("")
    const [alertMessage, setAlertMessage] = useState("")
    const [alertClassName, setAlertClassName] = useState("d-none")

    const navigate = useNavigate()

    const logOut = () => {
        setJwtToken("")
        navigate("/login")
    }

    useEffect(() => {
        if (jwtToken === "") {
            const requestObjects = {
                method: "GET",
                credentials: "include"
            }

            fetch(`${process.env.REACT_APP_BACKEND}/refresh`, requestObjects)
                .then((response) => response.json())
                .then((data) => {
                    if (data.access_token) {
                        setJwtToken(data.access_token)
                    }
                })
                .catch(error => {
                    console.log("user is not logged in", error)
                })
        }
    }, [jwtToken]);

    return (
        <div className={"container"}>
            <div className={"row"}>
                <div className="col">
                    <h1 className={"mt-3"}>Go Volley App</h1>
                </div>
                <div className="col text-end">
                    {jwtToken === ""
                        ? <Link to="/login"><span className={"badge bg-success"}>Login</span></Link>
                        : <a href="#!" onClick={logOut}><span className="badge bg-danger">Logout</span></a>
                    }
                </div>
                <hr className={"mb-3"}/>
            </div>
            <div className="row">
                <div className="col-md-2">
                    <nav>
                        <div className="list-group">
                            <Link to="/" className={"list-group-item list-group-item-action"}>Home</Link>
                            <Link to="/seasons" className={"list-group-item list-group-item-action"}>Seasons</Link>
                            <Link to="/teams" className={"list-group-item list-group-item-action"}>Teams</Link>
                            <Link to="/schedule" className={"list-group-item list-group-item-action"}>Schedule</Link>
                            {jwtToken !== "" &&
                                <>
                                    <Link to="/admin" className={"list-group-item list-group-item-action"}>Admin</Link>
                                    <Link to="/graphql"
                                          className={"list-group-item list-group-item-action"}>GraphQL</Link>
                                </>
                            }
                        </div>
                    </nav>
                </div>
                <div className="col-md-10">
                    <Alert message={alertMessage} className={alertClassName}/>
                    <Outlet context={{
                        jwtToken,
                        setJwtToken,
                        setAlertClassName,
                        setAlertMessage
                    }}/>
                </div>
            </div>
        </div>
    );
}

export default App;
