defmodule Solution do
  @spec first_uniq_char(s :: String.t()) :: integer
  def first_uniq_char(s) do
    l = to_charlist(s)
    f = Enum.frequencies(l)
    Enum.find_value(Enum.with_index(l), -1, fn {e, i} -> f[e] == 1 && i end)
  end
end
