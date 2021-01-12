import env from "../environment";

const get = async (url, options) =>
	await send(url, { method: "GET", ...options });

const post = async (url, body, options) =>
	await send(url, {
		method: "POST",
		body: JSON.stringify(body),
		headers: {
			"Content-Type": "application/json",
		},
		...options,
	});

const send = async (url, options) => {
	const dest = env.apiUrl + url;

	try {
		const res = await fetch(dest, options);

		switch (res.status) {
			case 200:
				const data = await res.json();
				return {
					ok: true,
					data: data,
				};
			default:
				try {
					const error = await res.json();
					return {
						ok: false,
						error: error.error,
					};
				} catch (e) {
					console.error(
						`Failed to read response from ${dest} (status: ${res.status})`,
						e
					);
					return {
						ok: false,
						error: e.toString() || res.statusText,
					};
				}
		}
	} catch (e) {
		console.error(
			`Failed to make a ${options.method} request to ${dest}`,
			e
		);
		return {
			ok: false,
			error: e.toString(),
		};
	}
};

export { get, post };
