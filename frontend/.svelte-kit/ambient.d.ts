
// this file is generated — do not edit it


/// <reference types="@sveltejs/kit" />

/**
 * Environment variables [loaded by Vite](https://vitejs.dev/guide/env-and-mode.html#env-files) from `.env` files and `process.env`. Like [`$env/dynamic/private`](https://kit.svelte.dev/docs/modules#$env-dynamic-private), this module cannot be imported into client-side code. This module only includes variables that _do not_ begin with [`config.kit.env.publicPrefix`](https://kit.svelte.dev/docs/configuration#env) _and do_ start with [`config.kit.env.privatePrefix`](https://kit.svelte.dev/docs/configuration#env) (if configured).
 * 
 * _Unlike_ [`$env/dynamic/private`](https://kit.svelte.dev/docs/modules#$env-dynamic-private), the values exported from this module are statically injected into your bundle at build time, enabling optimisations like dead code elimination.
 * 
 * ```ts
 * import { API_KEY } from '$env/static/private';
 * ```
 * 
 * Note that all environment variables referenced in your code should be declared (for example in an `.env` file), even if they don't have a value until the app is deployed:
 * 
 * ```
 * MY_FEATURE_FLAG=""
 * ```
 * 
 * You can override `.env` values from the command line like so:
 * 
 * ```bash
 * MY_FEATURE_FLAG="enabled" npm run dev
 * ```
 */
declare module '$env/static/private' {
	export const NVM_INC: string;
	export const rvm_use_flag: string;
	export const rvm_bin_path: string;
	export const TERM_PROGRAM: string;
	export const NODE: string;
	export const NVM_CD_FLAGS: string;
	export const rvm_quiet_flag: string;
	export const TERM: string;
	export const rvm_gemstone_url: string;
	export const SHELL: string;
	export const TMPDIR: string;
	export const rvm_docs_type: string;
	export const TERM_PROGRAM_VERSION: string;
	export const ORIGINAL_XDG_CURRENT_DESKTOP: string;
	export const MallocNanoZone: string;
	export const ZDOTDIR: string;
	export const rvm_hook: string;
	export const npm_config_local_prefix: string;
	export const PNPM_HOME: string;
	export const USER: string;
	export const NVM_DIR: string;
	export const rvm_gemstone_package_file: string;
	export const COMMAND_MODE: string;
	export const rvm_path: string;
	export const SSH_AUTH_SOCK: string;
	export const __CF_USER_TEXT_ENCODING: string;
	export const npm_execpath: string;
	export const rvm_proxy: string;
	export const rvm_ruby_file: string;
	export const rvm_prefix: string;
	export const rvm_silent_flag: string;
	export const rvm_ruby_make: string;
	export const PATH: string;
	export const _: string;
	export const npm_package_json: string;
	export const __CFBundleIdentifier: string;
	export const USER_ZDOTDIR: string;
	export const PWD: string;
	export const JAVA_HOME: string;
	export const npm_lifecycle_event: string;
	export const rvm_sdk: string;
	export const npm_package_name: string;
	export const LANG: string;
	export const VSCODE_GIT_ASKPASS_EXTRA_ARGS: string;
	export const XPC_FLAGS: string;
	export const npm_package_version: string;
	export const XPC_SERVICE_NAME: string;
	export const VSCODE_INJECTION: string;
	export const rvm_version: string;
	export const rvm_pretty_print_flag: string;
	export const SHLVL: string;
	export const HOME: string;
	export const rvm_script_name: string;
	export const VSCODE_GIT_ASKPASS_MAIN: string;
	export const rvm_ruby_mode: string;
	export const LOGNAME: string;
	export const rvm_alias_expanded: string;
	export const VSCODE_GIT_IPC_HANDLE: string;
	export const NVM_BIN: string;
	export const BUN_INSTALL: string;
	export const npm_config_user_agent: string;
	export const rvm_nightly_flag: string;
	export const VSCODE_GIT_ASKPASS_NODE: string;
	export const rvm_ruby_make_install: string;
	export const GIT_ASKPASS: string;
	export const rvm_niceness: string;
	export const rvm_ruby_bits: string;
	export const rvm_bin_flag: string;
	export const rvm_only_path_flag: string;
	export const COLORTERM: string;
	export const npm_node_execpath: string;
	export const NODE_ENV: string;
}

/**
 * Similar to [`$env/static/private`](https://kit.svelte.dev/docs/modules#$env-static-private), except that it only includes environment variables that begin with [`config.kit.env.publicPrefix`](https://kit.svelte.dev/docs/configuration#env) (which defaults to `PUBLIC_`), and can therefore safely be exposed to client-side code.
 * 
 * Values are replaced statically at build time.
 * 
 * ```ts
 * import { PUBLIC_BASE_URL } from '$env/static/public';
 * ```
 */
declare module '$env/static/public' {
	
}

/**
 * This module provides access to runtime environment variables, as defined by the platform you're running on. For example if you're using [`adapter-node`](https://github.com/sveltejs/kit/tree/master/packages/adapter-node) (or running [`vite preview`](https://kit.svelte.dev/docs/cli)), this is equivalent to `process.env`. This module only includes variables that _do not_ begin with [`config.kit.env.publicPrefix`](https://kit.svelte.dev/docs/configuration#env) _and do_ start with [`config.kit.env.privatePrefix`](https://kit.svelte.dev/docs/configuration#env) (if configured).
 * 
 * This module cannot be imported into client-side code.
 * 
 * ```ts
 * import { env } from '$env/dynamic/private';
 * console.log(env.DEPLOYMENT_SPECIFIC_VARIABLE);
 * ```
 * 
 * > In `dev`, `$env/dynamic` always includes environment variables from `.env`. In `prod`, this behavior will depend on your adapter.
 */
declare module '$env/dynamic/private' {
	export const env: {
		NVM_INC: string;
		rvm_use_flag: string;
		rvm_bin_path: string;
		TERM_PROGRAM: string;
		NODE: string;
		NVM_CD_FLAGS: string;
		rvm_quiet_flag: string;
		TERM: string;
		rvm_gemstone_url: string;
		SHELL: string;
		TMPDIR: string;
		rvm_docs_type: string;
		TERM_PROGRAM_VERSION: string;
		ORIGINAL_XDG_CURRENT_DESKTOP: string;
		MallocNanoZone: string;
		ZDOTDIR: string;
		rvm_hook: string;
		npm_config_local_prefix: string;
		PNPM_HOME: string;
		USER: string;
		NVM_DIR: string;
		rvm_gemstone_package_file: string;
		COMMAND_MODE: string;
		rvm_path: string;
		SSH_AUTH_SOCK: string;
		__CF_USER_TEXT_ENCODING: string;
		npm_execpath: string;
		rvm_proxy: string;
		rvm_ruby_file: string;
		rvm_prefix: string;
		rvm_silent_flag: string;
		rvm_ruby_make: string;
		PATH: string;
		_: string;
		npm_package_json: string;
		__CFBundleIdentifier: string;
		USER_ZDOTDIR: string;
		PWD: string;
		JAVA_HOME: string;
		npm_lifecycle_event: string;
		rvm_sdk: string;
		npm_package_name: string;
		LANG: string;
		VSCODE_GIT_ASKPASS_EXTRA_ARGS: string;
		XPC_FLAGS: string;
		npm_package_version: string;
		XPC_SERVICE_NAME: string;
		VSCODE_INJECTION: string;
		rvm_version: string;
		rvm_pretty_print_flag: string;
		SHLVL: string;
		HOME: string;
		rvm_script_name: string;
		VSCODE_GIT_ASKPASS_MAIN: string;
		rvm_ruby_mode: string;
		LOGNAME: string;
		rvm_alias_expanded: string;
		VSCODE_GIT_IPC_HANDLE: string;
		NVM_BIN: string;
		BUN_INSTALL: string;
		npm_config_user_agent: string;
		rvm_nightly_flag: string;
		VSCODE_GIT_ASKPASS_NODE: string;
		rvm_ruby_make_install: string;
		GIT_ASKPASS: string;
		rvm_niceness: string;
		rvm_ruby_bits: string;
		rvm_bin_flag: string;
		rvm_only_path_flag: string;
		COLORTERM: string;
		npm_node_execpath: string;
		NODE_ENV: string;
		[key: `PUBLIC_${string}`]: undefined;
		[key: `${string}`]: string | undefined;
	}
}

/**
 * Similar to [`$env/dynamic/private`](https://kit.svelte.dev/docs/modules#$env-dynamic-private), but only includes variables that begin with [`config.kit.env.publicPrefix`](https://kit.svelte.dev/docs/configuration#env) (which defaults to `PUBLIC_`), and can therefore safely be exposed to client-side code.
 * 
 * Note that public dynamic environment variables must all be sent from the server to the client, causing larger network requests — when possible, use `$env/static/public` instead.
 * 
 * ```ts
 * import { env } from '$env/dynamic/public';
 * console.log(env.PUBLIC_DEPLOYMENT_SPECIFIC_VARIABLE);
 * ```
 */
declare module '$env/dynamic/public' {
	export const env: {
		[key: `PUBLIC_${string}`]: string | undefined;
	}
}
