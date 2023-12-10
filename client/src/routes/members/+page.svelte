<script lang="ts">
  import { goto } from "$app/navigation";
  import { onMount } from "svelte";
  import { calculateAge } from "$lib/utils";
  import type { Member } from "$lib/types";

  let members: Member[];
  let showTable: boolean = true;
  let showTree: boolean = false;
  let error: { message: any };
  const memberAttributes = [
    "Firstname",
    "Middlename",
    "Lastname",
    "Age",
    "Gender",
    "City",
    "Country",
  ];

  onMount(() => {
    getMembers()
      .then((data) => {
        members = data;
      })
      .catch((err) => {
        error = err;
      });
  });

  async function getMembers() {
    const response = await fetch("http://localhost:8000/v1/members");
    const json = await response.json();
    return json.members;
  }

  function handleTableVisibility() {
    showTable = true;
    showTree = false;
  }

  function handleTreeVisibility() {
    showTable = false;
    showTree = true;
  }
</script>

<div class="flex align-stretch h-15 w-15 border-solid border-black m-4">
  <button on:click={handleTableVisibility} class="mr-4" type="button">
    <img src="/table-list-svgrepo-com.svg" alt="table" class="h-10 w-10" />
  </button>

  <button on:click={handleTreeVisibility} type="button">
    <img src="/family-tree-svgrepo-com.svg" alt="family tree" class="h-8 w-10" />
  </button>
</div>

{#if showTable && members}
  <div class="flex flex-col">
    <div class="overflow-x-auto sm:-mx-6 lg:-mx-8">
      <div class="inline-block min-w-full py-2 sm:px-6 lg:px-8">
        <div class="overflow-hidden">
          <table class="min-w-full text-left text-sm font-light">
            <thead class="border-b font-medium dark:border-neutral-500">
              <tr>
                {#each memberAttributes as heading}
                  <th class="px-6 py-4">{heading}</th>
                {/each}
              </tr>
            </thead>
            <tbody>
              {#each members as member}
                <tr
                  on:click={() => goto(`/members/${member.id}`)}
                  class="cursor-pointer border-b transition duration-300 ease-in-out hover:bg-neutral-100 dark:border-neutral-500 dark:hover:bg-neutral-600"
                >
                  {#each memberAttributes as attribute}
                    <td class="whitespace-nowrap px-6 py-4">
                      {#if attribute === "Age"}
                        {calculateAge(member.birthdate)}
                      {:else}
                        {member[attribute.toLowerCase()]}
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
{:else if error}
  <p style="color: red">{error.message}</p>
{:else if showTable && !members}
  <p>Loading...</p>
{/if}

{#if showTree && members}
  <p>Tree...</p>
{/if}
