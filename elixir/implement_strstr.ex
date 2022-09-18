defmodule Solution do
  @spec str_str(haystack :: String.t(), needle :: String.t()) :: integer
  def str_str(haystack, needle) do
    str_str(haystack, needle, 0)
  end

  defp str_str(haystack, needle, i) do
    if String.starts_with?(haystack, needle) do
      i
    else
      case haystack do
        <<_::utf8, rest::binary>> -> str_str(rest, needle, i + 1)
        _ -> -1
      end
    end
  end
end
