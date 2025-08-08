import http from 'k6/http';
import { sleep } from 'k6';

export default function () {
	// 发一个 HTTP 请求
	http.get('https://test.k6.io');
	console.log("hello")
	// 暂停 1 秒，模拟用户思考时间
	sleep(1);
}
