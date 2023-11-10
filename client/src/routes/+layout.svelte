<nav>
	<a href="/">Home</a>
	<a href="/about">About</a>
	<a href="/table">Table</a>
	<a class="nav-right-side" href="/login">Login</a>
</nav>


<style>

	nav {
		background-color: #333; 
		padding: 10px; 
		display: flex; /* Use flexbox to control positioning */
		justify-content: space-between; /* Push elements to opposite ends */
	}

	nav a {
		text-decoration: none; /* Remove underline from links */
		color: #fff; 
		margin: 0 10px; 
	}

	nav a:hover {
		text-decoration: underline; 
	}

	.nav-right-side {
		margin-left: auto;
	}

</style>

<slot />