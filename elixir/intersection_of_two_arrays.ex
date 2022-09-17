defmodule Solution do
  @spec intersect(nums1 :: [integer], nums2 :: [integer]) :: [integer]
  def intersect(nums1, nums2) do
    f1 = Enum.frequencies(nums1)
    f2 = Enum.frequencies(nums2)

    Enum.reduce(f1, [], fn {n, c}, acc -> List.duplicate(n, min(c, Map.get(f2, n, 0))) ++ acc end)
  end
end
