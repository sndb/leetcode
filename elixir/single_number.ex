defmodule Solution do
  @spec single_number(nums :: [integer]) :: integer
  def single_number(nums) do
    single_number(nums, %{})
  end

  defp single_number([], seen) do
    Enum.find_value(seen, fn {n, c} -> c == 1 && n end)
  end

  defp single_number([head | tail], seen) do
    single_number(tail, Map.update(seen, head, 1, &(&1 + 1)))
  end
end
