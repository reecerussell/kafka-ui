import React from "react";
import PropTypes from "prop-types";

const ruleSets = [
	{ rule: /^true|false,?$/, className: "boolean" },
	{ rule: /^[0-9]+,?$/, className: "number" },
	{ rule: /^"[^"]+",?$/, className: "string" },
	{ rule: /^null,?$/, className: "null" },
];

const Json = ({ data }) => {
	const keys = Object.keys(data);
	for (let i = 0; i < keys.length; i++) {
		const key = keys[i];
		try {
			data[key] = JSON.parse(data[key]);
		} catch {}
	}

	const json = JSON.stringify(data, null, 4);

	const renderLine = (line, idx) => {
		const test = line.match(/"[^"]*":/);
		if (!test) {
			return (
				<li key={idx}>
					<span className="line-number">{idx + 1}:</span> {line}
				</li>
			);
		}

		const match = test[0];
		const format = line.substring(0, line.indexOf(match));
		const propName = match.substring(1, match.length - 2);
		let value = line.substring(line.indexOf(match) + match.length + 1);

		for (let i = 0; i < ruleSets.length; i++) {
			const { rule, className } = ruleSets[i];
			if (!rule.test(value)) {
				continue;
			}

			if (value[value.length - 1] === ",") {
				value = (
					<>
						<span className={className}>
							{value.substring(0, value.length - 1)}
						</span>
						,
					</>
				);
			} else {
				value = <span className={className}>{value}</span>;
			}

			break;
		}

		return (
			<li key={idx}>
				<span className="line-number">{idx + 1}:</span> {format}"
				<span className="prop-name">{propName}</span>": {value}
			</li>
		);
	};

	return (
		<pre className="kui-json">
			<ul>{json.split("\n").map(renderLine)}</ul>
		</pre>
	);
};

Json.propTypes = {
	data: PropTypes.oneOfType([
		PropTypes.array.isRequired,
		PropTypes.object.isRequired,
	]),
};

export default Json;
