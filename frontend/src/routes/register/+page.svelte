<script>
    import { onMount } from 'svelte';

    let name = "";
    let surname = "";
    let email = "";
    let phone = "";
    let id;

    let camps = []

    function convertUnixTime(unixTime) {
        const date = new Date(unixTime * 1000);
        const day = date.getDate();
        const month = date.getMonth() + 1;
        const year = date.getFullYear();

        return `${day}.${month}. ${year}`;
    }

    function submitRegister() {
        fetch('/api/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                name: name,
                surname: surname,
                email: email,
                phone:  phone,
                camp_id: parseInt(id)

            })
        })
        .then(response => response.json())
    }

    onMount(async () => {
        fetch('/api/camp').then(response => response.json()).then(data => {
            camps = data.camps
        })
    })
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

<body class="overflow-hidden w-full h-screen">
    <h1 class="font-mono text-6xl text-white text-center mt-16">Registrace</h1>
    <div class="flex justify-center w-full">
        <div class="bg-gray-950 mt-20 flex justify-center flex-col p-10 rounded-xl gap-1 w-96">
            <label for="name" class="text-gray-300 font-mono">celé_jméno:</label>
            <div class="flex gap-1">
                <input type="text" placeholder="jméno" class=" font-mono w-40 h-10 rounded-xl p-2 text-center bg-slate-950 hover:bg-indigo-800 border border-indigo-500 border-2 text-gray-300 transition-all" bind:value={name}/>
                <input type="text" placeholder="příjmení" class="font-mono w-[155px] h-10 rounded-xl p-2 text-center bg-slate-950 hover:bg-indigo-800 border border-indigo-500 border-2 text-gray-300 transition-all" bind:value={surname}/>
            </div>
            <label for="email" class="text-gray-300 font-mono">e_mail:</label>
            <input type="email" class="font-mono w-80 h-10 rounded-xl p-2 text-center bg-slate-950 hover:bg-indigo-800 border border-indigo-500 border-2 text-gray-300 transition-all" bind:value={email}/>
            <label for="tel" class="text-gray-300 font-mono">telefon:</label>
            <input type="text" class="font-mono w-80 h-10 rounded-xl p-2 text-center bg-slate-950 hover:bg-indigo-800 border border-indigo-500 border-2 text-gray-300 transition-all" bind:value={phone}/>
        
            <label for="term" class="text-gray-300 mt-4 font-mono">termín:</label>
            <select bind:value={id} class="font-mono w-80 h-10 rounded-xl p-2 text-center bg-slate-950 hover:bg-indigo-800 border border-indigo-500 border-2 text-gray-300 transition-all">
                {#each camps as termin (termin)}
                <option value={termin.id}>{termin.name + " - " + convertUnixTime(termin.date)}</option>
                {/each}
            </select>
            <button on:click={submitRegister} class="font-mono w-40 h-12 bg-gray-950 mt-5 text-center rounded-xl self-center hover:bg-green-400 border border-green-400 border-2 transition-all text-gray-300 hover:text-gray-800">Registrovat</button>
        </div>
    </div>
</body>