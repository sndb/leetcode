#include <set>

using namespace std;

class SmallestInfiniteSet {
public:
	int popSmallest() {
		if (!elems.empty()) {
			int min = *elems.begin();
			if (min <= curr)
				elems.erase(min);
			if (min < curr)
				return min;
		}
		return curr++;
	}

	void addBack(int num) {
		if (num < curr)
			elems.insert(num);
	}

	int curr = 1;
	set<int> elems;
};
