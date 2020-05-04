import { clientObj } from '../client';
import { obj } from '../obj';

const mockedArg: string = 'foo'
const mockedProp: string ='foobar'

jest.spyOn(obj, 'prop').mockImplementation(() => mockedProp)

it('yurushite', () => {
  expect(clientObj(mockedArg)).toBe(`${mockedArg}_${mockedProp}`)
})
