import { Cs416Page } from './app.po';

describe('cs416 App', function() {
  let page: Cs416Page;

  beforeEach(() => {
    page = new Cs416Page();
  });

  it('should display message saying app works', () => {
    page.navigateTo();
    expect(page.getParagraphText()).toEqual('app works!');
  });
});
