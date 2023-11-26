<script>
    let islogged = false
    let username = "";
    let password = "";
</script>

<style>
	body {
		background-color: #e5e5f7;
		opacity: 0.83;
		background-image: radial-gradient(#444cf7 0.7px, rgb(2, 2, 15) 0.7px);
		background-size: 20px 20px;	

	}
</style>

<body class="flex w-full h-screen justify-center">
    {#if islogged}
        <div class="bg-gray-950 mt-40 mb-64 p-10 rounded-xl">

        </div>
    {:else}
    <div class="bg-gray-950 mt-40 mb-64 flex justify-center flex-col p-10 rounded-xl gap-1">
        <h1 class="text-gray-300 text-5xl font-mono text-center mt-[-20px]">Admin</h1>
        <label for="username" class="text-gray-300 font-mono">username:</label>
        <input type="text" class="font-mono w-80 h-10 rounded-xl p-2 text-center bg-slate-950 hover:bg-indigo-800 border border-indigo-500 border-2 text-gray-300 transition-all" bind:value={username}/>
        <label for="password" class="text-gray-300 font-mono">password:</label>
        <input type="password" class="font-mono w-80 h-10 rounded-xl p-2 text-center bg-slate-950 hover:bg-indigo-800 border border-indigo-500 border-2 text-gray-300 transition-all" bind:value={password}/>
        <button on:click={submitLogin} class="font-mono w-40 h-12 bg-gray-950 mt-5 text-center rounded-xl self-center hover:bg-green-400 border border-green-400 border-2 transition-all text-gray-300 hover:text-gray-800">Login</button>
    </div>
    {/if}
</body>