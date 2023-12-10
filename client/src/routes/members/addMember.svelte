<script lang="ts">
  function addMember(event) {
    event.preventDefault();
    // Logic to handle member addition
  }
</script>

<div class="flex justify-center items-center h-screen">
  <div class="bg-white border-black shadow overflow-hidden sm:rounded-lg w-96 p-8">
    <form on:submit|preventDefault={addMember} method="POST" class="flex flex-col">
      <label for="name" class="mb-2">Name</label>
      <input type="text" name="name" required class="border-solid border-2 mb-4 rounded-lg p-2" />

      <label for="relation" class="mb-2">Relation</label>
      <input
        type="text"
        name="relation"
        required
        class="border-solid border-2 mb-4 rounded-lg p-2"
      />

      <label for="birthdate" class="mb-2">Birthdate</label>
      <input
        type="date"
        name="birthdate"
        required
        class="border-solid border-2 mb-4 rounded-lg p-2"
      />

      <button
        type="submit"
        class="inline-flex h-10 w-full items-center justify-center gap-2 rounded border border-slate-300 bg-white p-2 text-sm font-medium text-black outline-none focus:ring-2 focus:ring-[#333] focus:ring-offset-1 disabled:cursor-not-allowed disabled:opacity-60"
      >
        Add Member
      </button>
    </form>
  </div>
</div>
