import env from "../environment";

export class Listener {
	constructor() {
		if (!window["WebSocket"]) {
			throw new Error("WebSockets are not supported by the browser!");
		}

		this.handlers = new Map();
	}

	Start() {
		this.conn = new WebSocket(env.wsUrl);
		this.conn.onclose = (evt) => {
			console.log("WebSocket connection closed!");
		};

		this.conn.onmessage = (evt) => {
			try {
				console.log("Event", evt);
				const messages = evt.data.split("\n");

				for (let i = 0; i < messages.length; i++) {
					const { type, payload } = JSON.parse(messages[i]);

					const handler = this.handlers.get(type);
					if (handler) {
						handler(payload);
					}
				}
			} catch (e) {
				console.error("Failed to read WebSocket message", e);
			}
		};
	}

	Close() {
		this.conn.close();
	}

	Handle(type, callback) {
		this.handlers.set(type, callback);
	}
}
