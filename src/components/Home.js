import Player from './../images/vol.jpg'
import {Link} from "react-router-dom";

const Home = () => {
    return <>
        <div className="text-center">
            <h2>Welcome to Home page</h2>
            <hr/>
            <Link to="/seasons">
                <img src={Player} className="img-thumbnail" alt="player"/>
            </Link>
        </div>
    </>
}

export default Home