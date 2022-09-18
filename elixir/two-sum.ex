defmodule Solution do
  @spec two_sum(nums :: [integer], target :: integer) :: [integer]
  def two_sum(nums, target) do
    two_sum(Enum.with_index(nums), target, %{})
  end

  defp two_sum([{v, i} | tail], target, seen) do
    if ii = seen[target - v] do
      [i, ii]
    else
      two_sum(tail, target, Map.put(seen, v, i))
    end
  end
end
