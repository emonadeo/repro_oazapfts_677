import { getFoo } from "./api.gen.ts";

const foo = await getFoo({
	baseUrl: "http://localhost:8080"
});

console.table([{ type: typeof foo.data, data: foo.data }]);
