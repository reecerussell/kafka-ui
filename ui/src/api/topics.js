import { Api } from "../utils";
import * as Actions from "../redux/actions/topicActions";

export const fetchTopics = () => (dispatch) => {
	dispatch(Actions.getTopics());

	return Api.get("topics")
		.then((res) => {
			if (!res.ok) {
				throw new Error(res.error);
			}

			dispatch(Actions.getTopicsSuccess(res.data));
		})
		.catch((error) => dispatch(Actions.getTopicsError(error.toString())));
};
