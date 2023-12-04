<script>
    import { onMount } from "svelte";

    let islogged = false
    let username = ""
    let password = ""
    let meals = []
    let participants = []

    onMount(async () => {
        fetch('/api/admin/participants').then(response => {
            if (!response.ok) {
                islogged = false;
            }
            else {
                islogged = true;
                fetch('/api/admin/participants').then(res => res.json()).then(data => {
                    participants = data.participants
                })
                fetch('/api/admin/food').then(res => res.json()).then(data => {
                    meals = data.foods
                })
            }
        })
    })

    function submitLogin() {
        fetch('/api/admin/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                username: username,
                password: password
            })
        }).then(res => islogged = res.status == 200)
    }
</script>

<style>
    @keyframes bg-animation {
        0% {
            background-position: 0 0;
        }
        100% {
            background-position: 0 100vh;
        }
    }

    body {
        background-color: #e5e5f7;
        background-image: radial-gradient(#444cf7 1px, rgb(2, 2, 15) 1px);
        background-size: 40px 40px;
        animation: bg-animation 60s linear infinite;
    }
</style>

<body class="overflow-hidden flex w-full h-screen justify-center">
    {#if islogged}
        <div class="overflow-scroll bg-gray-950 mt-40 mb-64 p-10 rounded-xl">
            {#each participants as participant}
                <div class="flex full text-white">
                    <h1>{participant.name} {participant.surname}</h1>
                    <h1>{participant.email}</h1>
                    <h1>{participant.phone}</h1>
                    <h1>{meals.find(item => item.id === participant.food_id).name}</h1>
                </div>
            {/each}
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