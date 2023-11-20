<script>
 let meals = ["pizza", "kebab", "durum", "meal4", "meal5"]
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
    <div class="bg-gray-950 opacity-80 w-11/12 mt-20 p-10 h-screen rounded-xl">
        <div class="grid grid-cols-3 gap-12 mb-10 place-content-center"> 
            {#each meals as meal}
            <div class="h-[360px] basis-1/3 bg-white rounded-3xl">
                {meal}
            </div>
            {/each}
        </div>

    </div>
</body>