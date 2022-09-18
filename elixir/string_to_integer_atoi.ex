defmodule Solution do
  @numbers ["0", "1", "2", "3", "4", "5", "6", "7", "8", "9"]
  @spec my_atoi(s :: String.t()) :: integer
  def my_atoi(s) do
    {m, s} =
      case String.trim_leading(s) do
        "+" <> rest -> {1, rest}
        "-" <> rest -> {-1, rest}
        s -> {1, s}
      end

    ["0" | String.graphemes(s)]
    |> Enum.take_while(fn x -> x in @numbers end)
    |> Enum.join()
    |> String.to_integer()
    |> (&*/2).(m)
    |> min(2 ** 31 - 1)
    |> max(-2 ** 31)
  end
end
