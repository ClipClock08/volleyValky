import {useEffect, useState} from "react";
import {Link} from "react-router-dom";

const Teams = () => {
    const [teams, setTeams] = useState([])

    useEffect(() => {
        const requestOptions = {
            method: "GET",
            headers: {
                "Content-Type": "application/json"
            }
        }

        fetch(`${process.env.REACT_APP_BACKEND}/teams`, requestOptions)
            .then((response) => response.json())
            .then((data) => {
                setTeams(data)
            })
            .catch(error => {
                console.log(error)
            })
    }, []);
    return <>
        <div className="text-center">
            <h2>Команди</h2>
            <hr/>
        </div>
        <div className="row">
            {teams.map((m) => (
                <div className="col-lg-4" key={m.id}>
                    <div className="card">
                        <img src={m.logo} className="card-img-top" alt="..."/>
                        <div className="card-body">
                            <h5 className="card-title">{m.name}</h5>
                        </div>
                        <ul className="list-group list-group-flush">
                            <li className="list-group-item">Заснована: {m.since}</li>
                            <li className="list-group-item">Локація: {m.city_id}</li>
                        </ul>
                        <div className="card-body">
                            <Link to={`/teams/${m.id}`} className="card-link">Детальніше</Link>
                        </div>
                    </div>
                </div>
            ))}
        </div>
    </>
}

export default Teams