import Messages from "./views/messages";
import Settings from "./views/settings";

import TopicSubnav from "./components/topicSubnav";

const routes = [
	{
		path: "/messages",
		component: Messages,
		exact: true,
		subnav: TopicSubnav,
	},
	{
		path: "/settings",
		component: Settings,
		exact: true,
	},
];

export default routes;
