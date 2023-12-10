<script lang="ts">
  function saveUser() {}
</script>

<div class="flex justify-center items-center h-screen">
  <div class="bg-white border-black shadow overflow-hidden sm:rounded-lg w-96 p-8">
    <form on:submit|preventDefault={saveUser} method="POST" class="flex flex-col">
      <label for="username" class="mb-2">Username</label>
      <input
        type="text"
        name="username"
        required
        class="border-solid border-2 mb-4 rounded-lg p-2"
      />

      <label for="password" class="mb-2">Password</label>
      <input
        type="password"
        name="password"
        required
        class="border-solid border-2 mb-4 rounded-lg p-2"
      />

      <label for="email" class="mb-2">Email</label>
      <input type="email" name="email" required class="border-solid border-2 mb-4 rounded-lg p-2" />

      <button
        type="submit"
        class="inline-flex h-10 w-full items-center justify-center gap-2 rounded border border-slate-300
                bg-white p-2 text-sm font-medium text-black outline-none focus:ring-2 focus:ring-[#333] focus:ring-offset-1
                disabled:cursor-not-allowed disabled:opacity-60"
      >
        Sign in
      </button>
    </form>
  </div>
</div>
