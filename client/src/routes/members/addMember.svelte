<script lang="ts">
  import { onMount } from "svelte";

  let countries: any = [];
  let selectedPhoto;

  onMount(async () => {
    await getCountries().then((data) => {
      countries = data;
    });
  });

  async function getCountries() {
    const response = await fetch("https://restcountries.com/v3.1/all");
    const json = await response.json();
    return json;
  }

  async function addMember(event: Event) {
    const form = event.target as HTMLFormElement;
    const formData = new FormData(form);
    console.log(formData);
    try {
      const response = await fetch("http://localhost:8000/v1/members", {
        method: "POST",
        body: formData,
      });

      if (!response.ok) {
        throw new Error(`Error: ${response.statusText}`);
      }

      const result = await response.json();
      console.log(result);
    } catch (error) {
      console.error(error);
    }
  }
  function handlePhotoUpload(event: Event) {
    // Further processing for photo upload
  }
</script>

<div class="flex justify-center items-center">
  <div class="bg-white border-black shadow overflow-hidden sm:rounded-lg w-3/4 p-7">
    <form
      on:submit|preventDefault={addMember}
      method="POST"
      action="/add"
      class="grid grid-cols-3 gap-20"
    >
      <div>
        <label for="firstname" class="block mb-2">
          First name <span class="text-red-500">*</span>
        </label>
        <input
          type="text"
          name="firstname"
          required
          class="border-solid border-2 rounded-lg p-2 mb-4"
        />

        <label for="middlename" class="block mb-2">Middle name</label>
        <input type="text" name="middlename" class="border-solid border-2 rounded-lg p-2 mb-4" />

        <label for="lastname" class="block mb-2">
          Last name <span class="text-red-500">*</span>
        </label>
        <input type="text" name="lastname" class="border-solid border-2 rounded-lg p-2 mb-4" />

        <label for="birthdate" class="block mb-2">Birthdate</label>
        <input type="date" name="birthdate" class="border-solid border-2 rounded-lg p-2 mb-4" />
      </div>

      <div>
        <label for="gender" class="block mb-2">Gender <span class="text-red-500">*</span></label>
        <select name="gender" required class="border-solid border-2 rounded-lg p-2 mb-4">
          <option value="">Select Gender</option>
          <option value="Male">Male</option>
          <option value="Female">Female</option>
          <option value="Other">Other</option>
        </select>

        <label for="country" class="block mb-2">Country</label>
        <select name="country" required class="border-solid border-2 rounded-lg w-full p-2 mb-4">
          <option value="">Select Country</option>
          {#each countries as country}
            <option value={country.name.common}>{country.flag} {country.name.common}</option>
          {/each}
        </select>

        <label for="city" class="block mb-2">City</label>
        <input type="text" name="city" class="border-solid border-2 rounded-lg p-2 mb-4" />

        <label for="photo" class="block mb-2">Photo</label>
        <input
          type="file"
          name="photo"
          accept="image/*"
          class="border-solid border-2 rounded-lg p-2 mb-4"
          on:change={handlePhotoUpload}
        />
      </div>

      <div>
        <label for="description" class="mb-2">Description</label>
        <textarea name="description" rows="10" class="border-solid border-2 rounded-lg p-2 w-full"
        ></textarea>

        <button
          type="submit"
          class="mt-4 inline-flex h-10 w-full items-center justify-center gap-2 rounded border
          border-slate-300 bg-white p-2 text-sm font-medium text-black outline-none focus:ring-2
          focus:ring-[#333] focus:ring-offset-1 disabled:cursor-not-allowed disabled:opacity-60"
        >
          Add Member
        </button>
      </div>
    </form>
  </div>
</div>
