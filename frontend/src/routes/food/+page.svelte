<script>
 import sha256 from 'crypto-js/sha256';
 import { onMount } from 'svelte';

 let meals = [];
 let name = "";
 let surname = "";
 let nameSubmit = false;

 function submitName() {
    if(name != "" && name[0] == name[0].toUpperCase() && surname != "" && surname[0] == surname[0].toUpperCase()) {
        nameSubmit = !nameSubmit
    }
 }

 function sendMeal(mealname) {
    fetch('/api/order/' + sha256(name+surname), {
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
    for(let i = 0;i < data.foods.length;i++) {
        meals.push(data.foods[i])
    }
 })
})

</script>

<style>
	body {
		background-color: #e5e5f7;
		opacity: 0.83;
		background-image: radial-gradient(#444cf7 0.7px, rgb(2, 2, 15) 0.7px);
		background-size: 20px 20px;	

	}
</style>
~
<body class="w-full h-screen justify">
    <h1 class="font-mono text-6xl text-white text-center mt-10">Výběr jídla</h1>
    <div class="flex justify-center">
        {#if nameSubmit && name != "" && surname != ""}
        <div class="grid grid-cols-3 gap-12 mb-10 w-full p-20 place-content-center"> 
            {#each meals as meal}
            <button class="aspect-square basis-1/3 bg-red-400 rounded-3xl text-center" use:sendMeal>{meal.name}</button>
            {/each}
        </div>
        {:else}
        <div class="bg-gray-950 mt-40 mb-64 flex justify-center flex-col p-10 rounded-xl gap-1">
            <label for="name" class="text-gray-300 font-mono">jméno:</label>
            <input type="text" class="font-mono w-80 h-10 rounded-xl p-2 text-center bg-slate-950 hover:bg-indigo-800 border border-indigo-500 border-2 text-gray-300 transition-all" bind:value={name}/>
            <label for="surname" class="text-gray-300 font-mono">příjmení:</label>
            <input type="text" class="font-mono w-80 h-10 rounded-xl p-2 text-center bg-slate-950 hover:bg-indigo-800 border border-indigo-500 border-2 text-gray-300 transition-all" bind:value={surname}/>
            <button on:click={submitName} class="font-mono w-40 h-12 bg-gray-950 mt-5 text-center rounded-xl self-center hover:bg-green-400 border border-green-400 border-2 transition-all text-gray-300 hover:text-gray-800">Pokračovat</button>
        </div>
        {/if}
    </div>
</body>