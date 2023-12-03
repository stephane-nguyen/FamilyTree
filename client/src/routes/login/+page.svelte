<script lang="ts">
</script>

<div class="bg-tree_login bg-cover bg-no-repeat flex justify-center items-center h-screen">
    <div class="w-full max-w-xs">
        <form
            method="post"
            action="/login"
            class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4"
        >
            <div class="mb-4">
                <input
                    class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700
                     leading-tight focus:outline-none focus:shadow-outline"
                    type="text"
                    name="user"
                    placeholder="Username"
                />
            </div>
            <div>
                <input
                    class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3
                     leading-tight focus:outline-none focus:shadow-outline"
                    type="password"
                    name="password"
                    placeholder="Password"
                />
            </div>
            <div class="flex justify-center">
                <button
                    class="bg-zinc-400 hover:bg-zinc-500 text-white font-bold py-2 px-4 rounded
                     focus:outline-none focus:shadow-outline"
                    type="submit"
                >
                    Login
                </button>
            </div>
            <div class="flex justify-center mb-2">
                <a href="/new_account" class="underline underline-offset-1">
                    Forgot password?
                </a>
            </div>
            <div class="flex justify-center">
                <button
                    class="bg-zinc-400 hover:bg-zinc-500 text-white font-bold py-2 px-4 rounded
                     focus:outline-none focus:shadow-outline"
                    type="submit"
                >
                    Create a new account
                </button>
            </div>
        </form>
    </div>
</div>
