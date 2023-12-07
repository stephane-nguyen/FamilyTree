<script lang="ts">

  import { goto } from '$app/navigation';
  import { onMount } from 'svelte';
  import { calculateAge } from '$lib/utils';

  async function getPersons() {
    const response = await fetch('http://localhost:8000/v1/persons');
    const json = await response.json();
    return json.persons;
  }

  const personAttributes = ["Firstname", "Middlename", "Lastname", "Age", "Gender", "City", "Country"];

</script>

{#await getPersons()}
  <p>Loading...</p>
{:then persons}
  <div class="flex flex-col">
    <div class="overflow-x-auto sm:-mx-6 lg:-mx-8">
      <div class="inline-block min-w-full py-2 sm:px-6 lg:px-8">
        <div class="overflow-hidden">
          <table class="min-w-full text-left text-sm font-light">
            <thead class="border-b font-medium dark:border-neutral-500">
              <tr>
                {#each personAttributes as heading}
                  <th class="px-6 py-4">{heading}</th>
                {/each}
              </tr>
            </thead>
            <tbody>
              {#each persons as person}
                <tr on:click={() => goto(`/members/${person.id}`)} 
                  class="cursor-pointer border-b transition duration-300 ease-in-out hover:bg-neutral-100 dark:border-neutral-500 dark:hover:bg-neutral-600">
                {#each personAttributes as attribute}
                  <td class="whitespace-nowrap px-6 py-4">
                  {#if attribute === "Age"}
                    {calculateAge(person.birthdate)}
                  {:else}
                    {person[attribute.toLowerCase()]}
                  {/if}
                  </td>
                {/each}
                </tr>
              {/each}
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
{:catch error}
	<p style="color: red">{error.message}</p>
{/await}