import React from "react";
import {
	HashRouter as Router,
	Route,
	Redirect,
	Switch,
} from "react-router-dom";
import NavBar from "./components/navbar";
import "./scss/styles.scss";

import routes from "./routes";

const App = () => {
	return (
		<>
			<Router>
				<NavBar />

				{routes.map(
					(route, key) =>
						route.subnav && (
							<Route
								key={key}
								path={route.path}
								exact={route.exact}
								render={(props) => <route.subnav {...props} />}
							/>
						)
				)}

				<main className="container">
					<Switch>
						{routes.map((route, key) => (
							<Route
								key={key}
								path={route.path}
								exact={route.exact}
								render={(props) => (
									<route.component {...props} />
								)}
							/>
						))}
						<Redirect to="/messages" />
					</Switch>
				</main>
			</Router>
		</>
	);
};

export default App;
