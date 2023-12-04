<script>
 import sha256 from 'crypto-js/sha256';
 import { onMount } from 'svelte';

 let meals = [];
 let name = "";
 let surname = "";
 let nameSubmit = false;
 let camps = "";
 let id;

 function convertUnixTime(unixTime) {
        const date = new Date(unixTime * 1000);
        const day = date.getDate();
        const month = date.getMonth() + 1;
        const year = date.getFullYear();

        return `${day}.${month}. ${year}`;
    }

 function submitName() {
    if(name != "" && name[0] == name[0].toUpperCase() && surname != "" && surname[0] == surname[0].toUpperCase()) {
        nameSubmit = !nameSubmit
    }
 }

 function sendMeal(mealname) {
    fetch('/api/order/' + sha256(name+surname+id), {
        method: 'POST',
        headers: {
            'Content-Type' : 'application/json'
        },
        body: JSON.stringify({
            'id': meals.find(item => item.name === mealname).id,
            'name': mealname
        })
    })
 }

 onMount(async () => {
    fetch('/api/food').then(response => response.json()).then(data => {
        meals = data.foods
    })
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
~
<body class="overflow-hidden w-full h-screen">
    <h1 class="font-mono text-6xl text-white text-center mt-10">Výběr jídla</h1>
    <div class="flex justify-center">
        {#if nameSubmit && name != "" && surname != ""}
        <div class="overflow-scroll grid grid-cols-3 gap-12 mb-10 w-full p-20 place-content-center"> 
            {#each meals as meal (meal)}
            <button style="background-image: url({meal.image_path}); background-size: cover; background-position: center;" class="aspect-square basis-1/3 rounded-3xl text-center font-mono text-4xl text-white uppercase" on:click={() => sendMeal(meal.name)}>{meal.name}</button>
            {/each}
        </div>
        {:else}
        <div class="bg-gray-950 mt-40 mb-64 flex justify-center flex-col p-10 rounded-xl gap-1">
            <label for="name" class="text-gray-300 font-mono">jméno:</label>
            <input type="text" class="font-mono w-80 h-10 rounded-xl p-2 text-center bg-slate-950 hover:bg-indigo-800 border border-indigo-500 border-2 text-gray-300 transition-all" bind:value={name}/>
            <label for="surname" class="text-gray-300 font-mono">příjmení:</label>
            <input type="text" class="font-mono w-80 h-10 rounded-xl p-2 text-center bg-slate-950 hover:bg-indigo-800 border border-indigo-500 border-2 text-gray-300 transition-all" bind:value={surname}/>
            <label for="term" class="text-gray-300 mt-4 font-mono">termín:</label>
            <select bind:value={id} class="font-mono w-80 h-10 rounded-xl p-2 text-center bg-slate-950 hover:bg-indigo-800 border border-indigo-500 border-2 text-gray-300 transition-all">
                {#each camps as termin (termin)}
                <option value={termin.id}>{termin.name + " - " + convertUnixTime(termin.date)}</option>
                {/each}
            </select>
            <button on:click={submitName} class="font-mono w-40 h-12 bg-gray-950 mt-5 text-center rounded-xl self-center hover:bg-green-400 border border-green-400 border-2 transition-all text-gray-300 hover:text-gray-800">Pokračovat</button>
        </div>
        {/if}
    </div>
</body>