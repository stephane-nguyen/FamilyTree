<script lang="ts">
  import { goto } from "$app/navigation";
  import { calculateAge } from "$lib/utils";
  import type { Member } from "$lib/types";

  export let members: Member[];
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
</script>

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
                class="cursor-pointer border-b transition duration-300 ease-in-out hover:bg-zinc-300 dark:border-neutral-500"
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
