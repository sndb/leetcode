#include <queue>

using namespace std;

class SeatManager {
public:
	SeatManager(int n) {
		for (int i = 1; i <= n; i++)
			seats.push(i);
	}

	int reserve() {
		int min = seats.top();
		seats.pop();
		return min;
	}

	void unreserve(int seatNumber) {
		seats.push(seatNumber);
	}

	priority_queue<int, vector<int>, greater<int>> seats;
};
