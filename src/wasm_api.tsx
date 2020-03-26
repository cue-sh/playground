// WasmAPIImpl is the lightweight container for our Wasm API. The TypeScript
// application will call OnChange to register callbacks that are interested
// in when API methods are added, and then read from CUECompile etc. The Go code
// will set CUECompile and then call FireOnChange.
export class WasmAPIImpl {
	private callbacks = new Array<WasmAPICallback>();
	public readonly CUECompile: ((mode: string, input: string) => CUECompileResponse) | undefined;

	OnChange(f: WasmAPICallback): void {
		if (this.callbacks.indexOf(f) != -1) {
			return;
		}
		this.callbacks.push(f);
	}

	FireOnChange(): void {
		for (let c of this.callbacks) {
			c(this);
		}
	}
}
