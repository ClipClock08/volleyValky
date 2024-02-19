import {useEffect, useState} from "react";
import {useParams} from "react-router-dom";

const Team = () => {
    const [team, setTeam] = useState({})
    let {id} = useParams();

    useEffect(() => {
        let myTeam = {
            id: 1,
            name: "StM",
            since: "2021-01-01",
            logo: "",
            city_id: 3
        }
        setTeam(myTeam)
    }, [id]);

    return <>
        <div className="text-center">
            <h2>Команда {team.name}</h2>
            <hr/>
        </div>
    </>
}

export default Team