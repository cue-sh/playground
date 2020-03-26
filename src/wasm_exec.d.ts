// wasm_exec.d.ts is a type definition that wraps the wasm_exec.js file
// copied from the Go source tree. It also provides the type definitions
// for the API our Go program will export to window.WasmAPI

declare class Go {
	constructor();
	importObject: any;
	run(v: any): void;
}

declare type WasmAPICallback = (v: WasmAPI) => void;

declare interface WasmAPI {
	FireOnChange(): void;

	// TODO: for some reason we do not get a compiler error
	// when passing a function with no parameters as an argument.
	// Work out why this is.
	OnChange(callback: WasmAPICallback): void;
	readonly CUECompile: ((input: string, func: string, output: string, inputVal: string) => CUECompileResponse) | undefined;
}

declare interface CUECompileResponse {
	readonly value: string;
	readonly error: string;
}

declare interface Window {
	WasmAPI: WasmAPI;
}
