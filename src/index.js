import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import {createBrowserRouter, RouterProvider} from "react-router-dom";
import ErrorPage from "./components/ErrorPage";
import Home from "./components/Home";
import Seasons from "./components/Seasons";
import Teams from "./components/Teams";
import Schedule from "./components/Schedule";
import ManagePage from "./components/ManagePage";
import GraphQL from "./components/GraphQL";
import Login from "./components/Login";
import Season from "./components/Season";
import Team from "./components/Team";

const router = createBrowserRouter([
    {
        path: "/",
        element: <App/>,
        errorElement: <ErrorPage/>,
        children: [
            {index: true, element: <Home/>},
            {
                path: "/seasons",
                element: <Seasons/>
            },
            {
                path: "/seasons/:id",
                element: <Season/>
            },
            {
                path: "/teams",
                element: <Teams/>
            },
            {
                path: "/teams/:id",
                element: <Team/>
            },
            {
                path: "/schedule",
                element: <Schedule/>
            },
            {
                path: "/admin",
                element: <ManagePage/>
            },
            {
                path: "/graphql",
                element: <GraphQL/>
            },
            {
                path: "/login",
                element: <Login/>
            }
        ]
    },
])

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
    <React.StrictMode>
        <RouterProvider router={router}/>
    </React.StrictMode>
);

