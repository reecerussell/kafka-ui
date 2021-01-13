import React from "react";
import classNames from "classnames";
import { Link, useLocation } from "react-router-dom";
import nav from "../nav";

const NavBar = () => {
	const { pathname } = useLocation();

	return (
		<nav className="navbar navbar-expand-md fixed-top navbar-dark bg-dark">
			<div className="container-fluid pl-md-0">
				<Link className="navbar-brand" to="/">
					Kafka UI
				</Link>
				<button className="navbar-toggler p-0 border-0" type="button">
					<span className="navbar-toggler-icon"></span>
				</button>
				<div className="collapse navbar-collapse">
					<ul className="navbar-nav me-auto mb-2 mb-md-0">
						{nav.map((item, key) => {
							const classes = classNames("nav-link", {
								active: pathname === item.path,
							});

							return (
								<li className="nav-item" key={key}>
									<Link className={classes} to={item.path}>
										{item.text}
									</Link>
								</li>
							);
						})}
					</ul>
				</div>
			</div>
		</nav>
	);
};

export default NavBar;
